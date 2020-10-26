import React, { useState } from "react";
import axios from "axios";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";
import Cookies from 'js-cookie';
import {
  Button,
  Form,
  Grid,
  Header,
  Image,
  Message,
  Segment,
} from "semantic-ui-react";
let endpoint = "http://localhost:8080";

function Tasks() {
  // set state variables
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [login, setLogin] = useState(false);
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
      .post(endpoint + "/auth/signin", {
        username: username,
        password: password,
      })
      .then(
        function (response) {
          setLogin(true);
          console.log(response);
          console.log(
            "Successfully login, look at the cookie, you'll see the sent token"
          );
          console.log(Cookies.set("token"))
        },
        (error) => {
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
            <Image src="/logo.png" /> Register your account
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

              <Button color="teal" fluid size="large" onClick={submitLogin}>
                Login
              </Button>
            </Segment>
          </Form>
          <Message>
            Not have account? <Link to="/register">Register</Link>
          </Message>
        </Grid.Column>
      </Grid>
    </Router>
  );
}

export default Tasks;
