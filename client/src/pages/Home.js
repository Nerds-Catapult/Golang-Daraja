import React, { useState, useEffect } from 'react'

import { Loader, Cards, FormField } from '../components'

const Home = () => {
  const [loading, setLoading] = useState(true)
  const [allPosts, setAllPosts] = useState([null])
  return (
    <section className='max-w-7xl mx-auto'>
      <div>
        <h1 className="font-extrabold text-[#222328] text-[32px]">
            <q>The Community Showcase</q>
        </h1>
        <p className='mt-2 text-[#666e75] text-[14px] max-w[500px]'>
          Browse Through a collectin of imaginative and visually stunning images generated 
          by Dawnix-AI
        </p>
      </div>
    </section>
)}

export default Home
