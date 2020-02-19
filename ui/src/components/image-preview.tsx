import React from "react";
import styled from "styled-components";

type MacroImageProps = {
    macroViewUrl: string
}

const ImageStyled = styled.img`
   display: block;
   margin-left: auto;
   margin-right: auto;
   max-width:100%;
   max-height:100%;
`;

export default class ImagePreview extends React.Component<MacroImageProps> {
    render() {
        return (
            <ImageStyled src={this.props.macroViewUrl}/>
        );
    }
}
