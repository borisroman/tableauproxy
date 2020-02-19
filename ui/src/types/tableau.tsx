import {PersonalAccessToken} from "./personal-access-tokens";
import {getMacroParameters} from "../api/atlassian";
import {getMacroImageURL} from "../api/tableauproxy";

export type MacroParameterStorage = {
    data: string,
    siteId: string,
    viewId: string,
    personalAccessTokenUUID: string,
    imageStyle?: string,
}

export type MacroParameters = {
    personalAccessTokenSelected?: PersonalAccessToken
    tableauSiteSelected?: TableauSite
    tableauViewSelected?: TableauView
    imageStyle: string
}

export type ImageRequestParameters = {
    personalAccessTokenUUID: string,
    siteId: string,
    viewId: string,
}

export type TableauSite = {
    id: string
    name: string,
    contentUrl: string,
}

export type TableauView = {
    id: string
    name: string,
    contentUrl: string,
}

export type TableauSiteResponse = {
    sites: Array<TableauSite>
}

export type TableauViewResponse = {
    views: Array<TableauView>
}

export function getMacroImageURLFromMacroParameters(callback: (macroImageURL: string) => void) {
    getMacroParameters((macroParameters => {
        if (macroParameters.personalAccessTokenSelected?.uuid !== undefined && macroParameters.tableauSiteSelected?.id !== undefined && macroParameters.tableauViewSelected?.id !== undefined) {
            let imageRequestParameters: ImageRequestParameters = {
                personalAccessTokenUUID: macroParameters.personalAccessTokenSelected?.uuid,
                siteId: macroParameters.tableauSiteSelected?.id,
                viewId: macroParameters.tableauViewSelected?.id,
            };

            getMacroImageURL(imageRequestParameters, (success, imageURL) => {
                if (success && imageURL !== undefined) {
                    callback(imageURL);
                }
            });
        }
    }));
}