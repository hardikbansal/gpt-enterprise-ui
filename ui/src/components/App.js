import React, {useState, useEffect} from 'react';
import {useNavigate, BrowserRouter, Route, Routes} from 'react-router-dom'
import Login from './Login';
import Chat from './Chat';
import ProtectedRoute from "./ProtectedRoute";

function App() {
    return (
        <BrowserRouter>
            <Routes>
                <Route path='/login' element={<Login/>}/>
                <Route path='/' element={<ProtectedRoute><Chat/></ProtectedRoute>}/>
            </Routes>
        </BrowserRouter>
    );
}

export default App;