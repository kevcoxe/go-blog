import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css';



function App() {
  const [posts, setPosts] = useState([])
  const fetchPosts = async () => {
    try {
      const response = await axios.get('/api/posts')
      setPosts(response.data)
    } catch (err) {
      console.log('error: ', err)
    }
  }

  useEffect(() => {
    fetchPosts()
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        {posts.map((post) => {
          return (
            <div key={ post.Id }>
              <p>{ post.Title }</p>
            </div>
          )
        })}
      </header>
    </div>
  );
}

export default App;
