import React from "react";
import Page, {Grid, GridColumn} from "@atlaskit/page";
import MacroForm from "../components/macro-form";
import {PersonalAccessToken} from "../types/personal-access-tokens";
import {getMacroImageURL, getPersonalAccessTokens, getTableauSites, getTableauViews} from "../api/tableauproxy";
import ImagePreview from "../components/image-preview";
import {ImageRequestParameters, MacroParameters, TableauSite, TableauView} from "../types/tableau";
import {getMacroParameters, setTriggerSaveMacroParameters} from "../api/atlassian";

type MacroEditViewProps = {}
type MacroEditViewState = {
    personalAccessTokens: Array<PersonalAccessToken>
    personalAccessTokenIsLoading: boolean
    personalAccessTokenSelected?: PersonalAccessToken
    tableauSites: Array<TableauSite>
    tableauSiteIsLoading: boolean
    tableauSiteSelected?: TableauSite
    tableauViews: Array<TableauView>
    tableauViewIsLoading: boolean
    tableauViewSelected?: TableauView
    imageStyle: string
    macroViewUrl: string
}

export default class MacroViewEdit extends React.Component<MacroEditViewProps, MacroEditViewState> {
    constructor(props: MacroEditViewProps) {
        super(props);
        this.state = {
            personalAccessTokens: [],
            personalAccessTokenIsLoading: true,
            personalAccessTokenSelected: undefined,
            tableauSites: [],
            tableauSiteIsLoading: true,
            tableauSiteSelected: undefined,
            tableauViews: [],
            tableauViewIsLoading: true,
            tableauViewSelected: undefined,
            imageStyle: "",
            macroViewUrl: "",
        };
    }

    componentDidMount() {
        setTriggerSaveMacroParameters(() => {
            let macroParameters: MacroParameters = {
                personalAccessTokenSelected: this.state.personalAccessTokenSelected,
                tableauSiteSelected: this.state.tableauSiteSelected,
                tableauViewSelected: this.state.tableauViewSelected,
                imageStyle: this.state.imageStyle,
            };

            return macroParameters;
        });

        this.getPersonalAccessTokens();

        getMacroParameters((macroParameters => {
            this.setState({
                personalAccessTokenSelected: macroParameters.personalAccessTokenSelected,
                tableauSiteSelected: macroParameters.tableauSiteSelected,
                tableauViewSelected: macroParameters.tableauViewSelected,
                imageStyle: macroParameters.imageStyle || "",
            }, () => {
                this.getSites();
                this.getViews();
                this.updateMacroViewUrl();
            });
        }));
    }

    updateMacroViewUrl() {
        if (this.state.personalAccessTokenSelected !== undefined && this.state.personalAccessTokenSelected?.uuid !== undefined && this.state.tableauSiteSelected !== undefined && this.state.tableauViewSelected !== undefined) {
            let imageRequestParameters: ImageRequestParameters = {
                personalAccessTokenUUID: this.state.personalAccessTokenSelected?.uuid,
                siteId: this.state.tableauSiteSelected.id,
                viewId: this.state.tableauViewSelected.id,
            };

            getMacroImageURL(imageRequestParameters, (success, imageURL) => {
                if (success && imageURL !== undefined) {
                    this.setState({macroViewUrl: imageURL});
                }
            });
        }
    }

    personalAccessTokenCallback(value?: PersonalAccessToken) {
        this.setState({personalAccessTokenSelected: value}, () => this.getSites());
    }

    tableauSiteCallback(value?: TableauSite) {
        this.setState({tableauSiteSelected: value}, () => this.getViews());
    }

    tableauViewCallback(value?: TableauView) {
        this.setState({tableauViewSelected: value}, () => this.updateMacroViewUrl());
    }

    imageStyleCallback(value: string) {
        this.setState({imageStyle: value});
    }

    getPersonalAccessTokens() {
        getPersonalAccessTokens(((success, personalAccessTokens) => {
            if (success && personalAccessTokens !== undefined && personalAccessTokens !== null) {
                this.setState({
                    personalAccessTokens: personalAccessTokens,
                    personalAccessTokenIsLoading: false
                });
            } else {
                this.setState({personalAccessTokens: [], personalAccessTokenIsLoading: false});
            }
        }));
    }

    getSites() {
        if (this.state.personalAccessTokenSelected !== undefined && this.state.personalAccessTokenSelected.uuid !== undefined) {
            getTableauSites(this.state.personalAccessTokenSelected.uuid, ((success, sites) => {
                if (success && sites !== undefined) {
                    this.setState({tableauSites: sites, tableauSiteIsLoading: false});
                }
            }));
        }
    }

    getViews() {
        if (this.state.personalAccessTokenSelected !== undefined && this.state.personalAccessTokenSelected.uuid !== undefined && this.state.tableauSiteSelected !== undefined) {
            getTableauViews(this.state.personalAccessTokenSelected.uuid, this.state.tableauSiteSelected?.id, ((success, views) => {
                if (success && views !== undefined) {
                    this.setState({tableauViews: views, tableauViewIsLoading: false});
                }
            }));
        }
    }

    render() {
        return (
            <div className="app">
                <div className="ac-content">
                    <Page>
                        <Grid>
                            <GridColumn medium={4}>
                                <MacroForm
                                    personalAccessTokens={this.state.personalAccessTokens}
                                    personalAccessTokenIsLoading={this.state.personalAccessTokenIsLoading}
                                    personalAccessTokenSelected={this.state.personalAccessTokenSelected}
                                    personalAccessTokenCallback={this.personalAccessTokenCallback.bind(this)}
                                    tableauSites={this.state.tableauSites}
                                    tableauSiteIsLoading={this.state.tableauSiteIsLoading}
                                    tableauSiteSelected={this.state.tableauSiteSelected}
                                    tableauSiteCallback={this.tableauSiteCallback.bind(this)}
                                    tableauViews={this.state.tableauViews}
                                    tableauViewIsLoading={this.state.tableauViewIsLoading}
                                    tableauViewSelected={this.state.tableauViewSelected}
                                    tableauViewCallback={this.tableauViewCallback.bind(this)}
                                    imageStyle={this.state.imageStyle}
                                    imageStyleCallback={this.imageStyleCallback.bind(this)}
                                />
                            </GridColumn>
                            <GridColumn medium={8}>
                                <ImagePreview macroViewUrl={this.state.macroViewUrl}/>
                            </GridColumn>
                        </Grid>
                    </Page>
                </div>
            </div>
        );
    }
}