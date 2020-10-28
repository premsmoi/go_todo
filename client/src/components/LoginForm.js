import React, { useState } from "react";
import Register from "./Register.js";
import axios from "axios";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import {
  Button,
  Form,
  Grid,
  Header,
  Image,
  Message,
  Segment,
} from "semantic-ui-react";
import "semantic-ui-css/semantic.min.css";
let endpoint = "http://127.0.0.1:8080";

function LoginForm() {
  // set state variables
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [login, setLogin] = useState(false);
  const [wrongPassActivated, setWrongPassActivated] = useState(false);
  // Functions
  function onChangeHandlerUsername(event) {
    setUsername(event.target.value);
    console.log(username);
  }

  function onChangeHandlerPassword(event) {
    setPassword(event.target.value);
    console.log(password);
  }

  function submitLogin(event) {
    console.log("Submit login");
    axios
      .post(
        endpoint + "/auth/signin",
        {
          username: username,
          password: password,
        },
        {
          withCredentials: true,
        }
      )
      .then(
        function (response) {
          setLogin(true);
          setWrongPassActivated(false);
          window.location.href = "/task";
        },
        (error) => {
          setWrongPassActivated(true);
          console.log(error);
        }
      );
  }

  return (
    <Router>
      <Grid
        textAlign="center"
        style={{ height: "100vh" }}
        verticalAlign="middle"
      >
        <Grid.Column style={{ maxWidth: 450 }}>
          <Header as="h2" color="teal" textAlign="center">
            <Image src="https://www.pinclipart.com/picdir/middle/4-49680_post-it-clipart-small-cartoon-post-it-note.png" />{" "}
            Login to see your todo list
          </Header>
          <Form size="large">
            <Segment stacked>
              <Form.Input
                fluid
                icon="user"
                iconPosition="left"
                placeholder="Userame"
                onChange={onChangeHandlerUsername}
              />
              <Form.Input
                fluid
                icon="lock"
                iconPosition="left"
                placeholder="Password"
                type="password"
                onChange={onChangeHandlerPassword}
              />
              <Link to={login && !wrongPassActivated ? "/task" : "/loginform"}>
                <Button fluid size="small" onClick={submitLogin}>
                  <p>Login</p>
                </Button>
              </Link>
            </Segment>
          </Form>
          <div>
            {wrongPassActivated && (
              <p style={{ color: "red" }}>Incorrect Username or password</p>
            )}
          </div>

          <Message>
            Don't have an account yet?{" "}
            <Link
              to="/register"
              onClick={() => {
                window.location.href = "/register";
              }}
            >
              Create account
            </Link>
          </Message>
          <Message size = {"mini"} color = {"teal"}>
            
          <Message size = {"small"} color = {"green"}>
              This is a small Todo-app project made with Go, React and MongoDB. Source code of this project is on <a href = "https://github.com/Generalkhun/go_todo">Github</a> :)
          </Message>
          <Image size = {"small"} floated = {"center"} src = "https://blog.mgechev.com/images/revive/revive.png"/>{" "}
          <Image size = {"mini"} src = "https://upload.wikimedia.org/wikipedia/commons/thumb/a/a7/React-icon.svg/220px-React-icon.svg.png"/>
          <Image size = {"mini"} src = "https://assets.stickpng.com/thumbs/58481021cef1014c0b5e494b.png"/>{" "}
         
                    </Message>
        </Grid.Column>
      </Grid>
      
      <Switch>
        <Route path="/register">
          <Register />
        </Route>
      </Switch>
    </Router>
  );
}

export default LoginForm;
