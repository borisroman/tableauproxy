export type PersonalAccessToken = {
    uuid?: string
    baseUrl?: string
    name?: string
    secret?: string
}
export type PersonalAccessTokenRow = {
    key: string,
    token: PersonalAccessToken,
    isNew: boolean
    needsSaving: boolean
    isSaving: boolean
    isDeleting: boolean
}
export type PersonalAccessTokensResponse = {
    personalAccessTokens: Array<PersonalAccessToken>
}

export function getPersonalAccessTokenRow(tokens: Array<PersonalAccessTokenRow>, key: string, callback: (personalAccessTokenRow?: PersonalAccessTokenRow) => void) {
    for (let i = 0; i < tokens.length; i++) {
        if (tokens[i].key === key) {
            callback(tokens[i]);
        }
    }
}

export function getPersonalAccessToken(tokens: Array<PersonalAccessToken>, uuid: string, callback: (personalAccessToken: PersonalAccessToken) => void) {
    for (let i = 0; i < tokens.length; i++) {
        if (tokens[i].uuid === uuid) {
            callback(tokens[i]);
        }
    }
}

export function updatePersonalAccessTokenRow(personalAccessTokenRows: Array<PersonalAccessTokenRow>, updatedPersonalAccessToken: PersonalAccessTokenRow, callback: (personalAccessTokenRows: Array<PersonalAccessTokenRow>) => void) {
    const newTokens = personalAccessTokenRows.map((token) => {
        if (token.key === updatedPersonalAccessToken.key) {
            return updatedPersonalAccessToken;
        }
        return token;
    });

    callback(newTokens);
}