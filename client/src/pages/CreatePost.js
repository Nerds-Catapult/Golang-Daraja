import React from 'react'
import { useNavigate } from 'react-router-dom'
import { preview } from '../assets/';
import { getRandomPrompt } from '../utils';
import { FormField, Loader } from '../components';
const CreatePost = () => {
  const handleSubmit = (e) => {
    e.preventDefault();
    // setLoading(true);
    // setTimeout(() => {
    //   setLoading(false);
    //   setGenerateImg(true);
    // }, 2000);
  };
  const navigate = useNavigate();
  const [form, setForm] = React.useState({
    name: '',
    prompt: '',
    image: '',
  });
  const [generateImg, setGenerateImg] = React.useState(false);
  const [loading, setLoading] = React.useState(false);
  return (
    <section className='max-w-7xl mx-auto'>
      <div>
        <h1 className="font-extrabold text-[#222328] text-[32px]">
          Create
        </h1>
        <p className='mt-2 text-[#666e75] text-[14px] max-w[500px]'>
          Create a collection of imaginative and visually stunning images generated
          by Dawnix-AI
        </p>
      </div>
      <form className="mt-16 max-w-3xl" onSubmit={handleSubmit}></form>
    </section>
    )
}

export default CreatePost
