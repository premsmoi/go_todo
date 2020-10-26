import React, { useState } from "react";
import axios from "axios";
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

function LoginForm() {
  // set state variables
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
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
    axios.post(endpoint + "/auth/signin", {
      username: username,
      password: password,
    }).then(res => (

    ));
  }

  return (
    <Grid textAlign="center" style={{ height: "100vh" }} verticalAlign="middle">
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

            <Button color="teal" fluid size="large" onClick={submitLogin}>
              Login
            </Button>
          </Segment>
        </Form>
        <Message>
          New to us? <a href="https://www.facebook.com/">Sign Up</a>
        </Message>
      </Grid.Column>
    </Grid>
  );
}

export default LoginForm;
