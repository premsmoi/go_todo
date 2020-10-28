import React, { useEffect, useState } from "react";
import axios from "axios";
import {
  Segment,
  Card,
  Header,
  Form,
  Input,
  Icon,
  Button,
} from "semantic-ui-react";
import "semantic-ui-css/semantic.min.css";
let endpoint = "http://127.0.0.1:8080";

function Tasks(props) {
  // state variable
  const [task, setTask] = useState("");
  const [item, setItem] = useState([]);
  const [username, setUsername] = useState("");

  //initiate tasks from the user
  GetTask();

  //Handler functions
  function onChangeHandlers(event) {
    setTask(event.target.value);
  }

  //Authen function
  function RefreshToken() {
    axios
      .get(endpoint + "/auth/refresh", { withCredentials: true })
      .then(() => {
        console.log("token refreshed");
      });
  }
  function logOut() {
    axios.get(endpoint + "/auth/logout", { withCredentials: true }).then(() => {
      console.log("loged out");
      window.location.href = "/loginform";
    });
  }

  //task functions
  // gettask: get all task from db
  function GetTask() {
    RefreshToken(); //refresh token
    useEffect(() => {
      axios
        .get(endpoint + "/task/getTasks", { withCredentials: true })
        .then((response) => {
          var responseData = response.data;
          // set username to display as welcome message
          setUsername(
            responseData.slice(
              responseData.indexOf("?") + 1,
              responseData.length - 1
            )
          );
          console.log(username);
          var todoArray = responseData.slice(0, responseData.indexOf("?") - 1);

          todoArray = JSON.parse(todoArray);
          if (todoArray) {
            var mappedItem = todoArray.map((item) => {
              let color = "yellow";
              if (item.status) {
                color = "green";
              }
              return (
                <Card key={item._id} color={color} fluid>
                  <Card.Content>
                    <Card.Header textAlign="left">
                      <div style={{ wordWrap: "break-word" }}>{item.task}</div>
                    </Card.Header>

                    <Card.Meta textAlign="right">
                      <Icon
                        name="check circle"
                        color="green"
                        onClick={() => completeTask(item._id)}
                      />
                      <span style={{ paddingRight: 10 }}>Done</span>
                      <Icon
                        name="undo"
                        color="yellow"
                        onClick={() => undoTask(item._id)}
                      />
                      <span style={{ paddingRight: 10 }}>Undo</span>
                      <Icon
                        name="delete"
                        color="red"
                        onClick={() => deleteTask(item._id)}
                      />
                      <span style={{ paddingRight: 10 }}>Delete</span>
                    </Card.Meta>
                  </Card.Content>
                </Card>
              );
            });
            setItem(mappedItem);
          } else {
            setItem([]);
          }
        });
    }, []);
  }

  function updateTask() {
    RefreshToken(); //refresh token
    axios
      .get(endpoint + "/task/getTasks", { withCredentials: true })
      .then((response) => {
        var responseData = response.data;
        // set username to display as welcome message

        setUsername(
          responseData.slice(
            responseData.indexOf("?") + 1,
            responseData.length - 1
          )
        );
        console.log(username);
        var todoArray = responseData.slice(0, responseData.indexOf("?") - 1);

        todoArray = JSON.parse(todoArray);
        if (todoArray) {
          var mappedItem = todoArray.map((item) => {
            let color = "yellow";
            if (item.status) {
              color = "green";
            }
            return (
              <Card key={item._id} color={color} fluid>
                <Card.Content>
                  <Card.Header textAlign="left">
                    <div style={{ wordWrap: "break-word" }}>{item.task}</div>
                  </Card.Header>

                  <Card.Meta textAlign="right">
                    <Icon
                      name="check circle"
                      color="green"
                      onClick={() => completeTask(item._id)}
                    />
                    <span style={{ paddingRight: 10 }}>Done</span>
                    <Icon
                      name="undo"
                      color="yellow"
                      onClick={() => undoTask(item._id)}
                    />
                    <span style={{ paddingRight: 10 }}>Undo</span>
                    <Icon
                      name="delete"
                      color="red"
                      onClick={() => deleteTask(item._id)}
                    />
                    <span style={{ paddingRight: 10 }}>Delete</span>
                  </Card.Meta>
                </Card.Content>
              </Card>
            );
          });
          setItem(mappedItem);
        } else {
          setItem([]);
        }
      });
  }

  function undoTask(id) {
    RefreshToken(); //refresh token
    axios
      .put(endpoint + "/task/undoTask/" + id, {}, { withCredentials: true })
      .then((res) => {
        console.log(res);
        updateTask();
      });
  }

  function completeTask(id) {
    RefreshToken(); //refresh token
    console.log("code is here");
    axios
      .put(endpoint + "/task/completeTask/" + id, {}, { withCredentials: true })
      .then((res) => {
        console.log("code is her2");
        console.log(res);
        updateTask();
      });
  }

  function deleteTask(id) {
    RefreshToken(); //refresh token
    //Becareful, axios.delete has different structure: header is on second argument
    axios
      .delete(endpoint + "/task/deleteTask/" + id, { withCredentials: true })
      .then((res) => {
        console.log(res);
        updateTask();
      });
  }

  function createTask(event) {
    RefreshToken(); //refresh token
    if (task !== "") {
      axios
        .post(
          endpoint + "/task/createTask",
          { task },
          { withCredentials: true }
        )
        .then((res) => {
          console.log("Task Added");
          updateTask();
        });
    }
  }

  //render
  return (
    <div>
      <Segment inverted basic = {true}>
        <div className="row">
          <Header className="header" as="h2" color = {"teal"}>
            Hi {username},
          </Header>
          <Header className="header" as="p" color = {"teal"}>
            Create some tasks ?
          </Header>
          
        </div>
      

      <div className="row">
        <Form>
          <Input
            type="text"
            name="task"
            onChange={onChangeHandlers}
            value={task}
            fluid
            placeholder="Create Task"
          />
          <Button color="teal" basic = {true} inverted onClick={createTask}>Create Task</Button>
        </Form>
      </div>
      </Segment>
      <div className="row">
        <Card.Group itemsPerRow={3}>{item}</Card.Group>
      </div>
      <div>
        <Button onClick={logOut}>Log out</Button>
      </div>
    </div>
  );
}

export default Tasks;
