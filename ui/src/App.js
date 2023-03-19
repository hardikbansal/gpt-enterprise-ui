import React, { useState, useEffect } from 'react';
import {useNavigate, BrowserRouter as Router, Route, Routes} from 'react-router-dom'
import Login from './components/Login';
import Chat from './components/Chat';

function App() {
    return (
      <Router>
        <Routes>
          <Route path='chat' element={<Chat/>} />
          <Route path='/' element={<Login/>} />
        </Routes>
      </Router>
    );
}
export default App;