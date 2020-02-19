import React from "react";
import styled from "styled-components";

type ImageProps = {
    macroViewUrl: string
    imageRef: React.RefObject<HTMLImageElement>
    imageDidLoad: () => void
}

const ImageStyled = styled.img`
   display: block;
   margin-left: auto;
   margin-right: auto;
   max-width:100%;
`;

export default class Image extends React.Component<ImageProps> {
    constructor(props: ImageProps) {
        super(props);
        this.state = {
            imageRef: React.createRef(),
        };
    }

    render() {
        return (
            <ImageStyled
                src={this.props.macroViewUrl}
                ref={this.props.imageRef}
                onLoad={this.props.imageDidLoad}
            />
        );
    }
}
