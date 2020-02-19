import React, {Component} from "react";
import Select from "@atlaskit/select";

import {Field} from "@atlaskit/form";
import {PersonalAccessToken} from "../types/personal-access-tokens";
import {TableauSite, TableauView} from "../types/tableau";
import TextArea from "@atlaskit/textarea";

type MacroFormProps = {
    personalAccessTokens: Array<PersonalAccessToken>
    personalAccessTokenIsLoading: boolean
    personalAccessTokenSelected?: PersonalAccessToken
    personalAccessTokenCallback: (value?: PersonalAccessToken) => void
    tableauSites: Array<TableauSite>
    tableauSiteIsLoading: boolean
    tableauSiteSelected?: TableauSite
    tableauSiteCallback: (value?: TableauSite) => void
    tableauViews: Array<TableauView>
    tableauViewIsLoading: boolean
    tableauViewSelected?: TableauView
    tableauViewCallback: (value?: TableauView) => void
    imageStyle: string
    imageStyleCallback: (value: string) => void
}

export default class MacroForm extends Component<MacroFormProps> {
    render() {
        return (
            <div>
                <Field<PersonalAccessToken>
                    name="personal-access-token"
                    label="Select a Personal Access Token"
                >
                    {() => (
                        <Select<PersonalAccessToken>
                            value={this.props.personalAccessTokenSelected}
                            options={this.props.personalAccessTokens}
                            isSearchable
                            getOptionLabel={(token): string => {
                                return token.name + " - " + token.baseUrl;
                            }}
                            onChange={(value) => this.props.personalAccessTokenCallback(value as PersonalAccessToken)}
                        />
                    )}
                </Field>
                <Field<TableauSite>
                    name="tableau-site"
                    label="Select a Tableau Site"
                >
                    {() => (
                        <Select<TableauSite>
                            value={this.props.tableauSiteSelected}
                            options={this.props.tableauSites}
                            isSearchable
                            getOptionLabel={(token): string => {
                                return token.name;
                            }}
                            onChange={(value) => this.props.tableauSiteCallback(value as TableauSite)}
                            isDisabled={this.props.personalAccessTokenSelected === undefined}
                        />
                    )}
                </Field>
                <Field<TableauView>
                    name="tableau-view"
                    label="Select a Tableau View"
                >
                    {() => (
                        <Select<TableauView>
                            value={this.props.tableauViewSelected}
                            options={this.props.tableauViews}
                            isSearchable
                            getOptionLabel={(token): string => {
                                return token.name;
                            }}
                            onChange={(value) => this.props.tableauViewCallback(value as TableauView)}
                            isDisabled={this.props.tableauSiteSelected === undefined}
                        />
                    )}
                </Field>
                <Field<string, HTMLTextAreaElement>
                    name="image-style"
                    label="Add custom CSS to the <img> tag (only used when exporting to PDF)"
                >
                    {() => (
                        <TextArea
                            value={this.props.imageStyle}
                            resize="smart"
                            name="image-style-textarea"
                            isCompact
                            isMonospaced
                            minimumRows={4}
                            onChange={(e) => this.props.imageStyleCallback(e.target.value)}
                        />
                    )}
                </Field>
            </div>
        );
    }
}