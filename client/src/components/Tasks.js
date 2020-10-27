import React, { Component, useEffect, useState } from "react";
import axios from "axios";
import { Card, Header, Form, Input, Icon, Button } from "semantic-ui-react";
import 'semantic-ui-css/semantic.min.css'
let endpoint = "http://127.0.0.1:8080";

function Tasks(props) {
  // state variable
  const [task, setTask] = useState("");
  const [item, setItem] = useState([]);

  //initiate tasks from the user
  GetTask();

  //Handler functions
  function onChangeHandlers(event) {
    setTask(event.target.value);
  }
  function onSubmitHandler() {
    console.log("pRINTING task", task);
    if (task) {
      axios
        .post(
          endpoint + "/task",
          {
            task,
          },
          {
            withCredentials: true,
          }
        )
        .then((res) => {
          GetTask();
          setTask("");
          console.log(res);
        });
    }
  }
  //task functions
  // gettask: get all task from db
  function GetTask() {
    console.log("hey");
    useEffect(() => {
      axios
        .get(endpoint + "/task/getTasks", { withCredentials: true })
        .then((response) => {
          // var res = JSON.stringify(response);
          // console.log(response)
          // console.log(response.data)
        
          var todoArray = JSON.parse(
            response.data.replace("null","") // remove at the end 
          )
          console.log(todoArray);
          console.log(typeof(todoArray))
          if (todoArray) {
            var mappedItem = (
              (todoArray).map((item) => {
                let color = "yellow";
                if (item.status) {
                  color = "green";
                }
                return (
                  <Card key={item._id} color={color} fluid>
                    <Card.Content>
                      <Card.Header textAlign="left">
                        <div style={{ wordWrap: "break-word" }}>
                          {item.task}
                        </div>
                      </Card.Header>

                      <Card.Meta textAlign="right">
                        <Icon
                          name="check circle"
                          color="green"
                          //onClick={() => this.updateTask(item._id)}
                        />
                        <span style={{ paddingRight: 10 }}>Done</span>
                        <Icon
                          name="undo"
                          color="yellow"
                          //onClick={() => this.undoTask(item._id)}
                        />
                        <span style={{ paddingRight: 10 }}>Undo</span>
                        <Icon
                          name="delete"
                          color="red"
                          //onClick={() => this.deleteTask(item._id)}
                        />
                        <span style={{ paddingRight: 10 }}>Delete</span>
                      </Card.Meta>
                    </Card.Content>
                  </Card>
                );
              })
            ) 
            setItem(mappedItem)
          } else {
            setItem([]);
          }
        });
    },[]);
  }

  return (
    <div>
      <div className="row">
        <Header className="header" as="h2">
          Your todolist
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
          <Button onClick={onSubmitHandler}>Create Task</Button>
        </Form>
      </div>
      <div className="row">
        <Card.Group>{item}</Card.Group>
      </div>
    </div>
  );
}

export default Tasks;