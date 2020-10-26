import logo from "./logo.svg";
import "./App.css";
import { Container } from "react-bootstrap";
import LoginForm from "./components/LoginForm";
import Register from "./components/Register";
import Tasks from "./components/Tasks";
import Homepage from "./components/Homepage.js";
import { BrowserRouter as Router, Redirect,Switch, Route, Link } from "react-router-dom";

function App() {
  return (
    <Router>
      <div>
        <Container fluid>
          <Switch>
            
            <Route path="/loginform">
              <LoginForm />
            </Route>
            <Route path="/register">
              <Register />
            </Route>
            <Route path="/task">
              <Tasks />
            </Route>
            <Route>
              <Homepage />
            </Route>
          </Switch>
        </Container>
      </div>
    </Router>
  );
}

export default App;
