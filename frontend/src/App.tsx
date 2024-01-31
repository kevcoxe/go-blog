
import { useEffect, useState } from 'react';
import axios from 'axios';


// post type definition
type Post = {
  ID: number
  Title: string
  Content: string
}

interface CreatePostProps {
  handleSubmit: (title: string, content: string) => Promise<Error|null>
}

function CreatePost({ handleSubmit }: CreatePostProps) {
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')

  const handleCreate = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault()

    const err = await handleSubmit(title, content)
    if (err) {
      console.error('Error creating post', err)
      return
    }

    setTitle('')
    setContent('')
  }

  return (
    <form className='flex flex-col gap-2 my-3 mx-24' onSubmit={handleCreate}>
      <h1 className='text-5xl'>Create Post</h1>
      <div className='flex flex-col gap-2'>
        <label htmlFor='title'>Title</label>
        <input name='title' value={title} type='text' id='title' className='p-2 rounded-lg text-black' onChange={(e) => setTitle(e.target.value)} />
      </div>
      <div className='flex flex-col gap-2'>
        <label htmlFor='content'>Content</label>
        <textarea name='content' id='content' className='p-2 rounded-lg text-black' value={content} onChange={(e) => setContent(e.target.value)} />
      </div>
      <button className='bg-slate-600 p-2 rounded-lg' type='submit'>Create</button>
    </form>
  )
}


interface PostsProps {
  posts: Post[]
  setPosts: (posts: Post[]) => void
}

function Posts({ posts, setPosts }: PostsProps) {

  const handleDelete = async (id: number) => {
    if (!id) return

    try {
      await axios.delete(`/api/posts/${id}`);
      setPosts(posts.filter(post => post.ID !== id))
    } catch (error) {
      console.error('Error deleting post', error)
    }
  }

  return (
    <div className='flex flex-col gap-2 m-3'>
      <h1 className='text-5xl'>Posts</h1>
      <ul className='grid grid-cols-2 gap-4 w-full'>
        {posts.map(post => (
          <li className='col-span-1 p-5 border-4 border-white rounded-lg flex flex-col w-full' key={`${post.Title}${post.ID}`}>
            <h2 className='text-3xl'>{post.Title}</h2>
            <p className='text-lg'>{post.Content}</p>
            <div className='flex justify-end'>
              <button className='bg-slate-600 p-2 rounded-lg'>Edit</button>
              <button className='bg-red-600 p-2 rounded-lg ml-2' onClick={() => handleDelete(post.ID)}>Delete</button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  )
}

function App() {

  const [posts, setPosts] = useState<Post[]>([])
  const [poll, setPoll] = useState(true)

  useEffect(() => {
    axios.get('/api/posts')
      .then(response => {
        setPosts(response.data)
      })
  }, [])

  // if poll is true, fetch posts every 5 seconds
  useEffect(() => {
    if (poll) {
      const interval = setInterval(async () => {
        const response = await axios.get('/api/posts')
        setPosts(response.data)
      }, 5000)
      return () => clearInterval(interval)
    }
  }, [poll])

  const createPost = async (title: string, content: string): Promise<Error|null> => {
    try {
      await axios.post('/api/posts', { title, content })
    } catch (error) {
      console.error('Error creating post', error)
      return Error.apply(error)
    }

    // fetch posts again
    const response = await axios.get('/api/posts')
    setPosts(response.data)

    return null
  }

  const buttonColor = poll ? 'bg-red-600' : 'bg-slate-600'

  return (
    <div className='flex bg-slate-800 w-screen h-screen text-white flex-col overflow-y-scroll'>
      <button className={`${buttonColor} p-2 rounded-lg m-3`} onClick={() => setPoll(!poll)}>{poll ? 'Stop' : 'Start'} Polling</button>
      <CreatePost handleSubmit={createPost}/>
      <Posts posts={posts} setPosts={setPosts} />
    </div>
  )
}

export default App
