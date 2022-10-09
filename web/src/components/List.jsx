import React from "react";
import Post from "./Post.jsx";

export function List({posts}) {
    return <div className="flex flex-col justify-center">
        {posts.length !== 0 && (
            <div className="flex mx-auto justify-start mt-4 max-w-lg">
                <h1 className="md:text-2xl text-2xl font-bold pt-px text-orange-700 ">
                    Posts
                </h1>
            </div>
        )}
        {posts.map((post) => {
            return <Post key={post.id} post={post} />;
        })}
    </div>
}