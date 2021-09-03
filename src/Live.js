import {LiveError, LivePreview, LiveProvider} from "react-live";
import React from "react";
import EventDetails from "./EventDetails";

class Live extends React.Component {

    render = () => {
        const scope = { EventDetails };

        const code = `<EventDetails />`;
        const fixture={};
        return <LiveProvider code={code} scope={scope} fixture={fixture}>
            {/*<LiveEditor />*/}
            <LiveError />
            <LivePreview />
        </LiveProvider>;
    }
}

export default Live;