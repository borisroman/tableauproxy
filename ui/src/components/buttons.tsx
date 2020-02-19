import React, {Component} from "react";
import CheckButton from "./button-confirm";
import RemoveButton from "./button-remove";


type ButtonsProps = {
    isNew: boolean,
    needsSaving: boolean,
    isSaving: boolean,
    isDeleting: boolean,
    createPersonalAccessTokenCallback: () => void,
    updatePersonalAccessTokenCallback: () => void,
    deletePersonalAccessTokenCallback: () => void,
}


export default class Buttons extends Component<ButtonsProps> {
    render() {
        if (this.props.isNew) {
            return (
                <CheckButton
                    onClick={() => this.props.createPersonalAccessTokenCallback()}
                    isSaving={this.props.isSaving}
                />
            );
        } else if (this.props.needsSaving) {
            return (
                <div>
                    <RemoveButton
                        onClick={() => this.props.deletePersonalAccessTokenCallback()}
                        isDeleting={this.props.isDeleting}
                    />
                    <CheckButton
                        onClick={() => this.props.updatePersonalAccessTokenCallback()}
                        isSaving={this.props.isSaving}
                    />
                </div>
            );
        } else {
            return (
                <RemoveButton
                    onClick={() => this.props.deletePersonalAccessTokenCallback()}
                    isDeleting={this.props.isDeleting}
                />
            );
        }
    }
}