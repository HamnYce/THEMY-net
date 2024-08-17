/** @type {import('next').NextConfig} */
const nextConfig = {
  env: {
    NEXT_PUBLIC_JSON_COLUMNS: process.env.NEXT_PUBLIC_JSON_COLUMNS,
    NEXT_PUBLIC_JSON_FILE_PATH: process.env.NEXT_PUBLIC_JSON_FILE_PATH,
    NEXT_PUBLIC_TEST_JSON_FILE_PATH: process.env.NEXT_PUBLIC_TEST_JSON_FILE_PATH,
    NEXT_PUBLIC_IS_TEST_MODE: process.env.NEXT_PUBLIC_IS_TEST_MODE,
    NEXT_PUBLIC_DEBUG_MODE: process.env.NEXT_PUBLIC_DEBUG_MODE,
  },
};

//DEBUG: To check if ENV variables are being read.
console.log('Next Config Env:', nextConfig.env);

export default nextConfig;
