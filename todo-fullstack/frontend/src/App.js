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
        .then(res => res.json())
        .then(setTodos);
    }
  }, [isLoggedIn]);

  const handleRegister = async e => {
    e.preventDefault();
    try {
      await register(username, password);
      setAuthMode("login");
      setError("");
      alert("Registration successful! Please login.");
    } catch {
      setError("Registration failed");
    }
  };

  const handleLogin = async e => {
    e.preventDefault();
    try {
      await login(username, password);
      setIsLoggedIn(true);
      setError("");
    } catch {
      setError("Login failed");
    }
  };

  const handleLogout = () => {
    logout();
    setIsLoggedIn(false);
    setTodos([]);
    setUsername("");
    setPassword("");
  };

  const addTodo = async e => {
    e.preventDefault();
    const res = await fetch(API, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${getToken()}`
      },
      body: JSON.stringify({ task, done: false })
    });
    const newTodo = await res.json();
    setTodos([...todos, newTodo]);
    setTask("");
  };

  const toggleTodo = async id => {
    const todo = todos.find(t => t.id === id);
    const res = await fetch(`${API}/${id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${getToken()}`
      },
      body: JSON.stringify({ ...todo, done: !todo.done })
    });
    const updated = await res.json();
    setTodos(todos.map(t => (t.id === id ? updated : t)));
  };

  const deleteTodo = async id => {
    await fetch(`${API}/${id}`, {
      method: "DELETE",
      headers: { Authorization: `Bearer ${getToken()}` }
    });
    setTodos(todos.filter(t => t.id !== id));
  };

  if (!isLoggedIn) {
    return (
      <div style={{ maxWidth: 400, margin: "auto" }}>
