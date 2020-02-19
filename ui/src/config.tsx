export const BASE_URL = (() => {
    if (process.env.REACT_APP_BASE_URL) {
        return process.env.REACT_APP_BASE_URL;
    }

    return "http://localhost:8080";
})();

export const DEVELOPMENT_JWT_TOKEN = ((): string => {
    if (process.env.REACT_APP_DEVELOPMENT_JWT_TOKEN) {
        return process.env.REACT_APP_DEVELOPMENT_JWT_TOKEN;
    }

    return "";
})();

export const DEVELOPMENT_MACRO_PARAMETERS = ((): string => {
    if (process.env.REACT_APP_DEVELOPMENT_MACRO_PARAMETERS) {
        return process.env.REACT_APP_DEVELOPMENT_MACRO_PARAMETERS;
    }

    return "";
})();
