import React, { useState } from "react";
import LoginForm from "./LoginForm.js";
import {Container} from "semantic-ui-react"
let endpoint = "http://localhost:8080";

function Firstpage() {
  return (
    <div>
      <Container>
        <LoginForm />
      </Container>
    </div>
  );
}

export default Firstpage;
