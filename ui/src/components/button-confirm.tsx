import React, {Component} from "react";
import Button from "@atlaskit/button";
import CheckIcon from "@atlaskit/icon/glyph/check";
import styled from "styled-components";
import {gridSize} from "@atlaskit/theme/constants";

type CheckButtonProps = {
    onClick: () => void
    isSaving: boolean
}

const ButtonWrapper = styled.div`
  display: inline-block;
  margin-left: ${gridSize}px;
  margin-right: ${gridSize}px;
`;

export default class CheckButton extends Component<CheckButtonProps> {
    render() {
        return (
            <ButtonWrapper>
                <Button
                    iconBefore={<CheckIcon label="check icon"/>}
                    onClick={() => this.props.onClick()}
                    appearance={"warning"}
                    isLoading={this.props.isSaving}
                >
                    Confirm
                </Button>
            </ButtonWrapper>
        );
    }
}