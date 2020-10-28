import React, { useState } from "react";
import axios from "axios";
import { BrowserRouter as Router} from "react-router-dom";

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

function Register() {
  // set state variables
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [invalidUsername, setInvalidUsername] = useState(false)
  const [nonBlank, setNonBlank] = useState(false)
  // Functions
  function onChangeHandlerUsername(event) {
    setUsername(event.target.value);
    if (username !== "" ){
      setNonBlank(true)
    }
    console.log(username);
  }

  function onChangeHandlerPassword(event) {
    setPassword(event.target.value);
    if (password !== "" ){
      setNonBlank(true)
    }
    console.log(password);
  }

  function submitReigister(event) {
    if (username === "" || password === ""){
      setInvalidUsername(true)
      return
    }
    console.log("Submit Register");
    axios
      .post(endpoint + "/register", {
        username: username,
        password: password,
      })
      .then(
        function (response) {
          console.log(response);
          console.log("Successfully Register");
          window.location.href = "/loginform";
        },
        (error) => {
          setInvalidUsername(true)
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
            <Image src="https://banner2.cleanpng.com/20180425/pxe/kisspng-colored-pencil-clip-art-5ae0902950fb02.3622818715246664093317.jpg" /> Create your account
          </Header>
          <Form size="large">
            <Segment stacked>
              <Form.Input
                fluid
                icon="user"
                iconPosition="left"
                placeholder="Create your username here..."
                onChange={onChangeHandlerUsername}
              />
              <Form.Input
                fluid
                icon="lock"
                iconPosition="left"
                placeholder="Create your password here..."
                type="password"
                onChange={onChangeHandlerPassword}
              />
              <div>
                {invalidUsername && (
                  <p style={{ color: "red" }}>{nonBlank ? "the username already being used": "Please fill both username and password"}</p>
                )}
              </div>

              <Button color="teal" fluid size="large" onClick={submitReigister}>
                Register
              </Button>
            </Segment>
          </Form>
        </Grid.Column>
      </Grid>
    </Router>
  );
}

export default Register;
