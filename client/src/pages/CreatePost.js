import React from 'react'
import { useNavigate } from 'react-router-dom'
import { preview } from '../assets/';
import { getRandomPrompt } from '../utils';
import { FormField, Loader } from '../components';
const CreatePost = () => {
  const handleSubmit = (e) => {
    
  };
  const handleChange = (e) => {
    
  }
  const handleSurpriseMe = () => { 

  }
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
      <form className="mt-16 max-w-3xl" onSubmit={handleSubmit}>
        <div className='flex flex-col gap-5'>
          <FormField
            placeholder='Enter your name'
            value={form.name}
            LabelName='name'
            type='text'
            name='name'
            handleChange={handleChange}
          />
          <FormField
            placeholder='A cat smiling'
            value={form.prompt}
            LabelName='prompt'
            type='text'
            name='prompt'
            handleChange={handleChange}
            isSurpriseMe
            handleSurpriseMe = {handleSurpriseMe}
          />
          <div className="relative bg-gray-50 border borser-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 w-64 p-3 h-64 flex justify-center items-center">
            {form.photo ? (
              <img src={form.photo}
                alt={form.prompt}
                className="w-full h-full object-contain" />
            ) : (<img src={preview} alt='Preview' className='w-9/12 h-9/12 object-contain opacity-40'/>)}
          </div>
        </div>
      </form>
    </section>
  )
}

export default CreatePost
