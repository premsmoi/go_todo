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
          setWrongPassActivated(false)
          window.location.href = "/task"
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
            <Image src="/logo.png" /> Login to see your todo list
          </Header>
          <Form size="large">
            <Segment stacked>
              <Form.Input
                fluid
                icon="user"
                iconPosition="left"
                placeholder="E-mail address"
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
              Register
            </Link>
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
