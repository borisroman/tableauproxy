import React from "react";
import {RouteComponentProps} from "react-router";
import Image from "../components/image";
import {resizeFrame} from "../api/atlassian";
import {getMacroImageURLFromMacroParameters} from "../types/tableau";


interface MacroViewProps extends RouteComponentProps<any> {
}

interface MacroViewState {
    macroViewUrl: string
    imageRef: React.RefObject<HTMLImageElement>,
    imageDidLoad: () => void
}


export default class MacroView extends React.Component<MacroViewProps, MacroViewState> {
    constructor(props: MacroViewProps) {
        super(props);
        this.state = {
            macroViewUrl: "",
            imageRef: React.createRef<HTMLImageElement>(),
            imageDidLoad: this.imageDidLoad
        };
    }

    imageDidLoad() {
        if (this.state.imageRef.current !== null) {
            resizeFrame("100%", this.state.imageRef.current.height + "px");
        }
    }

    componentDidMount() {
        resizeFrame("100%", "100%");

        getMacroImageURLFromMacroParameters((macroImageURL => {
            this.setState({macroViewUrl: macroImageURL});
        }));
    }

    render() {
        return (
            <div className="app">
                <div className="ac-content">
                    <Image
                        macroViewUrl={this.state.macroViewUrl}
                        imageRef={this.state.imageRef}
                        imageDidLoad={this.state.imageDidLoad.bind(this)}
                    />
                </div>
            </div>
        );
    }
}
