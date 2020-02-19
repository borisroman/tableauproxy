import React, {Component} from "react";
import InlineEdit from "@atlaskit/inline-edit";
import ReadViewContainer from "@atlaskit/inline-edit/dist/esm/styled/ReadViewContainer";
import Textfield from "@atlaskit/textfield";
import styled from "styled-components";
import {gridSize} from "@atlaskit/theme/constants";

type EditableFieldProps = {
    value: () => string
    placeholder: () => string
    onConfirm: (value: string) => void
}

const FieldWrapper = styled.div`
  margin-bottom: ${gridSize}px;
`;

export default class EditableField extends Component<EditableFieldProps> {
    render() {
        return (
            <FieldWrapper>
                <InlineEdit
                    onConfirm={this.props.onConfirm}
                    defaultValue={this.props.value()}
                    editView={(fieldProps) => (
                        <Textfield
                            placeholder={this.props.placeholder()}
                            {...fieldProps}
                            autoFocus
                        />
                    )}
                    readView={() => (
                        <ReadViewContainer>
                            {this.props.value() || this.props.placeholder()}
                        </ReadViewContainer>
                    )}
                    hideActionButtons
                    readViewFitContainerWidth
                />
            </FieldWrapper>
        );
    }
}