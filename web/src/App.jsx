import React, {useEffect, useState} from "react";
import Input from "../src/components/Input.jsx"
import {List} from "./components/List.jsx";
import {Login} from "./components/Login";
import {ListPost} from "./api/reddit.jsx";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

export default function App() {
    const [isLogined, setIsLogined] = useState(false)
    const [posts, setPosts] = useState([])
    useEffect( () => {
        let token = localStorage.getItem('token')
        if (token) {
            async function fetchData() {
                const res = await ListPost(token)
                setPosts(res.posts.reverse())
                setIsLogined(true)
            }
            fetchData()
        }
    }, [])
    return (
        <div className="App">
            <div>
                <ToastContainer />
            </div>
            {
                !isLogined ?
                    <Login  isLogined={isLogined} setIsLogined={setIsLogined} /> :
                    <div>
                        <Input setIsLogined={setIsLogined} setPosts={setPosts}/>
                        <List posts={posts}/>
                    </div>
            }
        </div>
    );
}