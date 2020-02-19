// Import AP from https://connect-cdn.atl-paas.net/all.js
import {DEVELOPMENT_JWT_TOKEN, DEVELOPMENT_MACRO_PARAMETERS} from "../config";
import {MacroParameters, MacroParameterStorage} from "../types/tableau";

declare var AP: any;

export function authenticate(callback: (jwt: string) => void) {
    // For debug purposes
    if (process.env.NODE_ENV === "production") {
        AP.context.getToken((jwt: string) => callback(jwt));
    } else if (process.env.NODE_ENV === "development") {
        callback(DEVELOPMENT_JWT_TOKEN);
    }
}

export function setTriggerSaveMacroParameters(macroParameters: () => MacroParameters) {
    // For debug purposes
    if (process.env.NODE_ENV === "production") {
        AP.dialog.getButton("submit").bind(function () {
            let macroParams = macroParameters();
            if (macroParams.tableauSiteSelected !== undefined && macroParams.tableauViewSelected !== undefined && macroParams.personalAccessTokenSelected?.uuid !== undefined) {
                let macroParameterStorage: MacroParameterStorage = {
                    data: JSON.stringify(macroParams),
                    siteId: macroParams.tableauSiteSelected.id,
                    viewId: macroParams.tableauViewSelected.id,
                    personalAccessTokenUUID: macroParams.personalAccessTokenSelected.uuid,
                    imageStyle: macroParams.imageStyle,
                };

                console.log(macroParameterStorage);
                AP.confluence.saveMacro(macroParameterStorage);
                AP.confluence.closeMacroEditor();
            }
            return true;
        });
    } else if (process.env.NODE_ENV === "development") {
        let macroParams = macroParameters();
        if (macroParams.tableauSiteSelected !== undefined && macroParams.tableauViewSelected !== undefined && macroParams.personalAccessTokenSelected?.uuid !== undefined) {
            let macroParameterStorage: MacroParameterStorage = {
                data: JSON.stringify(macroParams),
                siteId: macroParams.tableauSiteSelected.id,
                viewId: macroParams.tableauViewSelected.id,
                personalAccessTokenUUID: macroParams.personalAccessTokenSelected.uuid,
                imageStyle: macroParams.imageStyle,
            };

            console.log(macroParameterStorage);
            console.log("triggered the save");
        }
    }
}

export function getMacroParameters(callback: (macroParameters: MacroParameters) => void) {
    // For debug purposes
    if (process.env.NODE_ENV === "production") {
        AP.confluence.getMacroData(function (macroParams: MacroParameterStorage) {
            let macroParameters = JSON.parse(macroParams.data);
            console.log(macroParameters);
            callback(macroParameters);
        });
    } else if (process.env.NODE_ENV === "development") {
        let macroParameters = JSON.parse(DEVELOPMENT_MACRO_PARAMETERS);
        console.log(macroParameters);
        callback(macroParameters);
    }
}

export function resizeFrame(width: string, height: string) {
    // For debug purposes
    if (process.env.NODE_ENV === "production") {
        AP.resize(width, height);
    } else if (process.env.NODE_ENV === "development") {
        console.log("sized to parent");
    }
}

export function listenToAll() {
    AP.events.onAny((eventName: string) => {
        console.log(eventName);
    });
}