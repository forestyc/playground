import json
import argparse


def convert_to_markdown(postman_collection, name):
    markdown_output = []
    collection_info = postman_collection.get("info", {})
    collection_name = collection_info.get("name", "API Collection")
    if name is None:
        markdown_output.append(f"# {collection_name} API\n")
        get_markdown_env(postman_collection, markdown_output)
        get_markdown_api(postman_collection.get("item", None), markdown_output)
    else:
        markdown_output.append(f"# {name} API\n")
        get_markdown_env(postman_collection, markdown_output)
        target = get_target_item(postman_collection, name)
        get_markdown_api(target.get("item", None), markdown_output)
    return "\n".join(markdown_output)


def get_markdown_env(postman_collection, markdown_output):
    markdown_output.append("## Hosts\n")
    table_buf = "|Name|Value|\n"
    table_buf += "|:--|:--|\n"
    for env in postman_collection.get("variable", []):
        table_buf += "|" + env["key"] + "|" + env["value"] + "|\n"
    markdown_output.append(table_buf)


def get_target_item(postman_collection, name):
    for sub_item in postman_collection.get("item", []):
        if sub_item.get("name", None) == name:
            return sub_item
        else:
            target = get_target_item(sub_item, name)
            if target is not None:
                return target


def get_markdown_api(item, markdown_output, dept=1):
    if item is None:
        return
    for sub_item in item:
        get_markdown_item(sub_item, markdown_output, dept)
        get_markdown_api(sub_item.get("item", None), markdown_output, dept + 1)


def get_markdown_item(item, markdown_output, dept=2):
    if item is None:
        return
    name = item.get("name", "Unnamed Endpoint")
    method = item.get("request", {}).get("method", "GET")
    url = item.get("request", {}).get("url", {}).get("raw", None)
    headers = item.get("request", {}).get("header", [])
    body = item.get("request", {}).get("body", {}).get("raw", None)
    responses = item.get("response", [])
    sharp = ''
    for _ in range(dept):
        sharp += '#'
    markdown_output.append(f"{sharp} {name}\n")
    if url is None:
        return
    markdown_output.append(f"- **Method**: {method}")
    markdown_output.append(f"- **URL**: `{url}`")

    # Headers
    markdown_output.append("- **Headers**:")
    for header in headers:
        key = header.get("key", "Unknown")
        value = header.get("value", "Unknown")
        markdown_output.append(
            f"  - `{key}`: {value} (Required, **Env**)" if "{{" in value else f"  - `{key}`: {value}")

    # Body (if available)
    if body:
        markdown_output.append("\n- **Body**:\n")
        markdown_output.append(f"```json\n{body.strip()}\n```\n")

    # Responses
    if responses:
        markdown_output.append("- **Responses**:")
        for response in responses:
            code = response.get("code", 200)
            markdown_output.append(f"  - **Status**: {code}")
            body = response.get("body", "{}").strip()
            if body:
                markdown_output.append(f"  - **Body**: \n```json\n{body}\n```")

    markdown_output.append("\n---\n")


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='help')
    parser.add_argument('-i', '--input', type=str, help='JSON of collection v2.1')
    parser.add_argument('-d', '--directory', type=str, help='JSON of collection v2.1')
    parser.add_argument('-o', '--output', type=str, help='markdown name', default='api.md')
    args = parser.parse_args()

    # Load Postman Collection JSON
    with open(args.input, 'r', encoding='utf-8') as f:
        postman_collection = json.load(f)

    # Convert to Markdown
    markdown_output = convert_to_markdown(postman_collection, args.directory)

    # Output the markdown to a file
    with open(args.output, 'w', encoding='utf-8') as f:
        f.write(markdown_output)

    print(f"Conversion complete! Markdown saved as {args.output}")
