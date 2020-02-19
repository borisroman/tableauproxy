/// <reference types="react-scripts" />
declare namespace NodeJS {
    interface ProcessEnv {
        NODE_ENV: "development" | "production" | "test"
        PUBLIC_URL: string
        REACT_APP_DEVELOPMENT_JWT_TOKEN: string
        REACT_APP_DEVELOPMENT_MACRO_PARAMETERS: string
        REACT_APP_BASE_URL: string
    }
}