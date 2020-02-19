import React from "react";
import {BrowserRouter, Route, Switch} from "react-router-dom";
import AdminPersonalAccessTokens from "./pages/admin-personal-access-tokens";
import MacroViewEdit from "./pages/macro-view-edit";
import MacroView from "./pages/macro-view";

export default function Router() {
    return (
        <BrowserRouter>
            <Switch>
                <Route path="/ui/admin/personal-access-tokens" component={AdminPersonalAccessTokens}/>
                <Route path="/ui/macro/view" component={MacroView}/>
                <Route path="/ui/macro/edit-view" component={MacroViewEdit}/>
            </Switch>
        </BrowserRouter>
    );
}
