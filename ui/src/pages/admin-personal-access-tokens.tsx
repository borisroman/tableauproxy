import {
    PersonalAccessToken,
    PersonalAccessTokenRow,
    updatePersonalAccessTokenRow
} from "../types/personal-access-tokens";
import React, {Component} from "react";
import {v4 as uuidv4} from "uuid";
import {HeadType, RowType} from "@atlaskit/dynamic-table/types";
import EditableField from "../components/editable-field";
import Buttons from "../components/buttons";
import {
    createPersonalAccessToken,
    deletePersonalAccessToken,
    getPersonalAccessTokens,
    updatePersonalAccessToken
} from "../api/tableauproxy";
import Page from "@atlaskit/page";
import PageHeader from "@atlaskit/page-header";
import Button, {ButtonGroup} from "@atlaskit/button";
import NewFeature24Icon from "@atlaskit/icon-object/glyph/new-feature/24";
import DynamicTable from "@atlaskit/dynamic-table";

type AdminPersonalAccessTokensProps = {}
type AdminPersonalAccessTokensState = {
    personalAccessTokens: Array<PersonalAccessTokenRow>
    isLoading: boolean
}

export default class AdminPersonalAccessTokens extends Component<AdminPersonalAccessTokensProps, AdminPersonalAccessTokensState> {
    constructor(props: AdminPersonalAccessTokensProps) {
        super(props);
        this.state = {
            personalAccessTokens: [],
            isLoading: true
        };
    }

    private createPersonalAccessToken(personalAccessTokenRow: PersonalAccessTokenRow) {
        this.personalAccessTokenRowIsSaving(personalAccessTokenRow);

        createPersonalAccessToken(personalAccessTokenRow.token, (success, newPersonalAccessToken) => {
            if (success && newPersonalAccessToken !== undefined) {
                personalAccessTokenRow.token = newPersonalAccessToken;
                personalAccessTokenRow.isNew = false;
                personalAccessTokenRow.needsSaving = false;
                personalAccessTokenRow.isSaving = false;

                updatePersonalAccessTokenRow(this.state.personalAccessTokens, personalAccessTokenRow, (updatedPersonalAccessTokenRows) => {
                    this.setState({personalAccessTokens: updatedPersonalAccessTokenRows});
                });
            } else {
                console.log("something went wrong with creating the token");
            }
        });

    }

    private deletePersonalAccessToken(personalAccessTokenRow: PersonalAccessTokenRow) {
        this.personalAccessTokenRowIsDeleting(personalAccessTokenRow);

        deletePersonalAccessToken(personalAccessTokenRow.token, (success) => {
            if (success) {
                this.setState((state) => {
                    const newTokens = state.personalAccessTokens.filter((token) => token.key !== personalAccessTokenRow.key);
                    return {personalAccessTokens: newTokens};
                });
            } else {
                console.log("something went wrong with deleting the token");
            }
        });

    }

    private updatePersonalAccessToken(personalAccessTokenRow: PersonalAccessTokenRow) {
        this.personalAccessTokenRowIsSaving(personalAccessTokenRow);

        updatePersonalAccessToken(personalAccessTokenRow?.token, (success, newPersonalAccessToken) => {
            if (success && newPersonalAccessToken !== undefined) {
                personalAccessTokenRow.token = newPersonalAccessToken;
                personalAccessTokenRow.isNew = false;
                personalAccessTokenRow.needsSaving = false;
                personalAccessTokenRow.isSaving = false;

                updatePersonalAccessTokenRow(this.state.personalAccessTokens, personalAccessTokenRow, (updatedPersonalAccessTokenRows) => {
                    this.setState({personalAccessTokens: updatedPersonalAccessTokenRows});
                });
            } else {
                console.log("something went wrong with updating the token");
            }
        });
    }

    private updateStateOfPersonalAccessTokenRow(key: string, personalAccessToken: PersonalAccessToken) {
        this.setState((state) => {
            const newTokens = state.personalAccessTokens.map((token) => {
                    if (token.key === key) {
                        if (personalAccessToken.baseUrl !== undefined) {
                            token.token.baseUrl = personalAccessToken.baseUrl;
                        }
                        if (personalAccessToken.name !== undefined) {
                            token.token.name = personalAccessToken.name;
                        }
                        if (personalAccessToken.secret !== undefined) {
                            token.token.secret = personalAccessToken.secret;

                        }
                        token.needsSaving = true;
                        return token;
                    }
                    return token;
                }
            );

            return {personalAccessTokens: newTokens};
        });
    }

    private personalAccessTokenRowIsSaving(personalAccessTokenRow: PersonalAccessTokenRow) {
        personalAccessTokenRow.isSaving = true;

        updatePersonalAccessTokenRow(this.state.personalAccessTokens, personalAccessTokenRow, (updatedPersonalAccessTokenRows) => {
            this.setState({personalAccessTokens: updatedPersonalAccessTokenRows});
        });
    }

    private personalAccessTokenRowIsDeleting(personalAccessTokenRow: PersonalAccessTokenRow) {
        personalAccessTokenRow.isDeleting = true;

        updatePersonalAccessTokenRow(this.state.personalAccessTokens, personalAccessTokenRow, (updatedPersonalAccessTokenRows) => {
            this.setState({personalAccessTokens: updatedPersonalAccessTokenRows});
        });
    }

    addRow() {
        const tokens = this.state.personalAccessTokens.slice();

        const newTokens = tokens.concat({
            key: uuidv4(),
            token: {
                uuid: "",
                baseUrl: "",
                name: "",
                secret: ""
            },
            isNew: true,
            needsSaving: false,
            isSaving: false,
            isDeleting: false,
        });

        this.setState({personalAccessTokens: newTokens});
    }

    mapRows(): Array<RowType> {
        return this.state.personalAccessTokens.map((personalAccessToken) => ({
            key: personalAccessToken.key,
            cells: [
                {
                    key: `${personalAccessToken.key}-baseUrl`,
                    content: (
                        <EditableField
                            value={() => {
                                if (personalAccessToken.token.baseUrl !== undefined) {
                                    return personalAccessToken.token.baseUrl;
                                }
                                return "";
                            }}
                            onConfirm={(value) => this.updateStateOfPersonalAccessTokenRow(personalAccessToken.key, {baseUrl: value})}
                            placeholder={() => "Enter the Base URL here..."}
                        />
                    ),
                    colSpan: 2
                },
                {
                    key: `${personalAccessToken.key}-name`,
                    content: (
                        <EditableField
                            value={() => {
                                if (personalAccessToken.token.name !== undefined) {
                                    return personalAccessToken.token.name;
                                }
                                return "";
                            }}
                            onConfirm={(value) => this.updateStateOfPersonalAccessTokenRow(personalAccessToken.key, {name: value})}
                            placeholder={() => "Enter the Name here..."}
                        />
                    ),
                    colSpan: 1
                },
                {
                    key: `${personalAccessToken.key}-secret`,
                    content: (
                        <EditableField
                            value={() => ""}
                            onConfirm={(value) => this.updateStateOfPersonalAccessTokenRow(personalAccessToken.key, {secret: value})}
                            placeholder={() => "Sensitive - Write only"}
                        />
                    ),
                    colSpan: 1
                },
                {
                    key: `${personalAccessToken.key}-buttons`,
                    content: (<Buttons
                        isNew={personalAccessToken.isNew}
                        needsSaving={personalAccessToken.needsSaving}
                        isSaving={personalAccessToken.isSaving}
                        isDeleting={personalAccessToken.isDeleting}
                        createPersonalAccessTokenCallback={() => this.createPersonalAccessToken(personalAccessToken)}
                        updatePersonalAccessTokenCallback={() => this.updatePersonalAccessToken(personalAccessToken)}
                        deletePersonalAccessTokenCallback={() => this.deletePersonalAccessToken(personalAccessToken)}
                    />),
                    colSpan: 1
                }
            ]
        }));
    }

    componentDidMount() {
        getPersonalAccessTokens(((success, personalAccessTokens) => {
            if (success && personalAccessTokens !== undefined && personalAccessTokens !== null) {
                const personalAccessTokenRows = personalAccessTokens.map(token => {
                    return {
                        key: uuidv4(),
                        token: token,
                        isNew: false,
                        needsSaving: false,
                        isSaving: false,
                        isDeleting: false,
                    };
                });

                this.setState({personalAccessTokens: personalAccessTokenRows, isLoading: false});
            } else {
                this.setState({personalAccessTokens: [], isLoading: false});
            }
        }));
    }

    render() {
        const head = {
            cells: [
                {
                    key: "baseUrl",
                    content: "Base URL",
                    isSortable: true,
                    colSpan: 2,
                },
                {
                    key: "personalAccessTokenName",
                    content: "Name",
                    isSortable: true,
                    colSpan: 1
                },
                {
                    key: "personalAccessTokenSecret",
                    content: "Secret",
                    colSpan: 1
                },
            ],
        };
        return (
            <div className="app">
                <div className="ac-content" style={{margin: 16}}>
                    <Page>
                        <PageHeader
                            actions={(
                                <ButtonGroup>
                                    <Button
                                        iconBefore={<NewFeature24Icon label="add icon"/>}
                                        onClick={() => {
                                            this.addRow();
                                        }}
                                    >
                                        Add Personal Access Token
                                    </Button>
                                </ButtonGroup>
                            )}
                        >
                            Personal Access Tokens
                        </PageHeader>
                        <DynamicTable
                            head={((): HeadType => {
                                if (this.state.personalAccessTokens.length === 0) {
                                    return {cells: []};
                                }
                                return head;
                            })()}
                            emptyView={<h2>No Personal Access Tokens are present</h2>}
                            rows={this.mapRows()}
                            rowsPerPage={10}
                            defaultPage={1}
                            loadingSpinnerSize="large"
                            isLoading={this.state.isLoading}
                            isFixedSize
                        />
                    </Page>
                </div>
            </div>
        );
    }
}