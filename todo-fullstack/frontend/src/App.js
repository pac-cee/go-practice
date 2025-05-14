import React, { useState, useEffect } from "react";
import { register, login, logout, getToken } from "./auth";

const API = "http://localhost:8080/todos";

function App() {
  const [todos, setTodos] = useState([]);
  const [task, setTask] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [authMode, setAuthMode] = useState("login");
  const [error, setError] = useState("");
  const [isLoggedIn, setIsLoggedIn] = useState(!!getToken());

  useEffect(() => {
    if (isLoggedIn) {
      fetch(API, {
        headers: { Authorization: `Bearer ${getToken()}` }
      })
