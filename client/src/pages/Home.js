import React, { useState, useEffect } from 'react'

import { Loader, Cards, FormField } from '../components';

const RenderCards = ({ data, title }) => {
  if(data?.length > 0 ) return data.map((post) => <Cards key={post._id} {...post} />)
}

const Home = () => {
  const [loading, setLoading] = useState(false);
  const [allPosts, setAllPosts] = useState([null]);
  const [searchText, setSearchText] = useState('');
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
      <div className="mt-16">
        <FormField />

      </div>
      <div className="mt-10">
        {loading ? (
          <div className='flex justify-center items-center'>
            <Loader />
          </div>
        ) : (
          <>
              {searchText &&( <h2 className='font-medium text-[#666e75] text-xl mb-3'>
                Search Results for <span className='text-[#222328]'>{searchText}</span>
              </h2>
              )}
              <div className="grid lg:grid-cols-4 sm:grid-cols-3 xs:grid-cols-2 grid-cols-1 gap-3">

              </div>
          </>
        )}
      </div>
    </section>
  )
}

export default Home
