import React from "react";
import Button from "@atlaskit/button";
import TrashIcon from "@atlaskit/icon/glyph/trash";
import styled from "styled-components";
import {gridSize} from "@atlaskit/theme/constants";

type RemoveButtonProps = {
    onClick: () => void
    isDeleting: boolean
}

const ButtonWrapper = styled.div`
  display: inline-block;
  margin-left: ${gridSize}px;
  margin-right: ${gridSize}px;
`;

export default class RemoveButton extends React.Component<RemoveButtonProps> {
    render() {
        return (
            <ButtonWrapper>
                <Button
                    iconBefore={<TrashIcon label="trash icon"/>}
                    onClick={this.props.onClick}
                    isLoading={this.props.isDeleting}
                >
                    Remove
                </Button>
            </ButtonWrapper>
        );
    }
}