import React, { useState, useEffect } from 'react';
import {useNavigate, BrowserRouter as Router, Route, Routes} from 'react-router-dom'
import Login from './Login';
import Chat from './Chat';

function App() {
    return (
      <Router>
        <Routes>
          <Route path='chat' element={<Chat/>} />
          <Route path='/' element={<Chat/>} />
        </Routes>
      </Router>
    );
}
export default App;