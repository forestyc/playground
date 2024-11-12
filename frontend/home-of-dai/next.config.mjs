/** @type {import('next').NextConfig} */
const nextConfig = {
    async rewrites() {
        return [
            {
                source: '/api/:path*',
                destination: `${process.env.HOST}/:path*`
            },
        ]
    }
};

export default nextConfig;
