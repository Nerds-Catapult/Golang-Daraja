import React from 'react'
import { useNavigate } from 'react-router-dom'
import { preview } from '../assets/';
import { getRandomPrompt } from '../utils';
import { FormField, Loader } from '../components';
const CreatePost = () => {
  const navigate = useNavigate();
  const [form, setForm] = React.useState({
    name: '',
    prompt: '',
    image: '',
  });
  const [generateImg, setGenerateImg] = React.useState(false);
  const [loading, setLoading] = React.useState(false);
  return (
    <div>createPost</div>
  )
}

export default CreatePost
