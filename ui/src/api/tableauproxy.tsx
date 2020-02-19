import {PersonalAccessToken, PersonalAccessTokensResponse} from "../types/personal-access-tokens";
import {BASE_URL} from "../config";
import {authenticate} from "./atlassian";
import {
    ImageRequestParameters,
    TableauSite,
    TableauSiteResponse,
    TableauView,
    TableauViewResponse
} from "../types/tableau";

export function createPersonalAccessToken(personalAccessToken: PersonalAccessToken, callback: (success: boolean, personalAccessToken?: PersonalAccessToken) => void) {
    authenticate((jwt) => {
        fetch(BASE_URL + "/personal-access-token?jwt=" + jwt, {
            method: "POST",
            headers: {
                "Accept": "application/json",
                "Content-Type": "application/json",
            },
            body: JSON.stringify(personalAccessToken)
        }).then(response => {
            if (!response.ok) {
                callback(false);
            }

            response.json().then((body: PersonalAccessToken) => {
                callback(true, body);
            });
        });
    });
}

export function updatePersonalAccessToken(personalAccessToken: PersonalAccessToken, callback: (success: boolean, personalAccessToken?: PersonalAccessToken) => void) {
    authenticate((jwt) => {
        fetch(BASE_URL + "/personal-access-token?jwt=" + jwt, {
            method: "PATCH",
            headers: {
                "Accept": "application/json",
                "Content-Type": "application/json",
            },
            body: JSON.stringify(personalAccessToken)
        }).then(response => {
            if (!response.ok) {
                callback(false);
            }

            response.json().then((body: PersonalAccessToken) => {
                callback(true, body);
            });
        });
    });
}

export function deletePersonalAccessToken(personalAccessToken: PersonalAccessToken, callback: (success: boolean) => void) {
    authenticate((jwt) => {
        fetch(BASE_URL + "/personal-access-token?jwt=" + jwt, {
            method: "DELETE",
            headers: {
                "Accept": "application/json",
                "Content-Type": "application/json",
            },
            body: JSON.stringify(personalAccessToken)
        }).then(response => {
                if (!response.ok) {
                    callback(false);
                } else {
                    callback(true);
                }
            }
        );
    });
}

export function getPersonalAccessTokens(callback: (success: boolean, personalAccessTokens?: Array<PersonalAccessToken>) => void) {
    authenticate((jwt) => {
        fetch(BASE_URL + "/personal-access-token?jwt=" + jwt, {
            method: "GET",
            headers: {
                "Accept": "application/json",
                "Content-Type": "application/json",
            },
        }).then(response => {
                if (!response.ok) {
                    callback(false);
                }

                response.json().then((body: PersonalAccessTokensResponse) => {
                    callback(true, body.personalAccessTokens);
                });
            }
        );
    });
}

export function getMacroImageURL(imageRequestParameters: ImageRequestParameters, callback: (success: boolean, imageURL?: string) => void) {
    authenticate((jwt) => {
        callback(true, BASE_URL + "/macro-image.png?jwt=" + jwt + "&siteId=" + imageRequestParameters.siteId + "&viewId=" + imageRequestParameters.viewId + "&personalAccessTokenUUID=" + imageRequestParameters.personalAccessTokenUUID);
    });
}

export function getTableauSites(personalAccessTokenUUID: string, callback: (success: boolean, sites?: Array<TableauSite>) => void) {
    authenticate((jwt) => {
        fetch(BASE_URL + "/tableau/sites?jwt=" + jwt + "&personalAccessTokenUUID=" + personalAccessTokenUUID, {
            method: "GET",
            headers: {
                "Accept": "application/json",
                "Content-Type": "application/json",
            },
        }).then(response => {
                if (!response.ok) {
                    callback(false);
                }

                response.json().then((body: TableauSiteResponse) => {
                    callback(true, body.sites);
                });
            }
        );
    });
}

export function getTableauViews(personalAccessTokenUUID: string, siteId: string, callback: (success: boolean, sites?: Array<TableauView>) => void) {
    authenticate((jwt) => {
        fetch(BASE_URL + "/tableau/views?jwt=" + jwt + "&personalAccessTokenUUID=" + personalAccessTokenUUID + "&siteId=" + siteId, {
            method: "GET",
            headers: {
                "Accept": "application/json",
                "Content-Type": "application/json",
            },
        }).then(response => {
                if (!response.ok) {
                    callback(false);
                }

                response.json().then((body: TableauViewResponse) => {
                    callback(true, body.views);
                });
            }
        );
    });
}
