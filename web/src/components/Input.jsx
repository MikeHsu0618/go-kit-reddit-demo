import React, { useState } from "react";
import {CreatePost} from "../api/reddit.jsx";
import {toast} from "react-toastify";

function parseJwt (token) {
    let base64Url = token.split('.')[1];
    let base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    let jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
};

export default function Input({setIsLogined, setPosts}) {
    const [title, setTitle] = useState("");
    const [content, setContent] = useState("");
    const handlePost = async () => {
        let token = localStorage.getItem("token")
        let claim = parseJwt(token)
        const data = await CreatePost(title, content, token, claim.user_id)
        if (!data) return
        setPosts(prev => [data.post, ...prev])
        toast.success("Create Post Success")
    };
    const handleLogout = () => {
        localStorage.removeItem("token")
        setIsLogined(false)
        toast.success("Log out")
    }

    return (
        <div className="flex justify-center">
            <div className="flex mx-auto items-center justify-center shadow-lg mt-20 mb-4 max-w-lg">
                <div
                    className="w-full max-w-xl bg-white rounded-lg px-4 pt-2"
                >
                    <div className="flex flex-wrap -mx-3 mb-6">
                        <h2 className="px-4 pt-3 pb-2 text-gray-800 text-lg">
                            Add a new Post
                        </h2>
                        <div className="w-full md:w-full px-3 mb-2 mt-2">
              <textarea
                  className="bg-gray-100 rounded border border-gray-400 leading-normal pt-1 resize-none w-full h-10  px-3 font-medium placeholder-gray-700 focus:outline-none focus:bg-white"
                  name="body"
                  placeholder="Type Your title"
                  value={title}
                  onChange={(e) => setTitle(e.target.value)}
                  required
              ></textarea>
              <textarea
                  className="bg-gray-100 rounded border border-gray-400 leading-normal resize-none w-full h-20 py-2 px-3 font-medium placeholder-gray-700 focus:outline-none focus:bg-white"
                  name="body"
                  placeholder="Type Your Cotent"
                  value={content}
                  onChange={(e) => setContent(e.target.value)}
                  required
              ></textarea>
                        </div>
                        <div className="w-full md:w-1/2 px-3">
                            <div className="-mr-1">
                                <button
                                    onClick={handleLogout}
                                    type="submit"
                                    className="bg-white text-gray-700 font-medium py-1 px-4 border border-gray-400 rounded-lg tracking-wide mr-1 hover:bg-gray-100"
                                >
                                    Log Out
                                </button>
                            </div>
                        </div>
                        <div className="w-full md:w-1/2 text-end px-3">
                            <div className="-mr-1">
                                <button
                                    onClick={handlePost}
                                    type="submit"
                                    className="bg-white text-gray-700 font-medium py-1 px-4 border border-gray-400 rounded-lg tracking-wide mr-1 hover:bg-gray-100"
                                >
                                    Post
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}