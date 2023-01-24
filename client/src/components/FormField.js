import React from 'react'

const FormField = ({ LabelName, type, name, placeholder, value, handleChange, isSurpriseMe, handleSurpriseMe }) => {
  return (
    <div>
      <div className='flex items-center gap-2 mb-2'>
        <label
          htmlFor={name}
          className="block text-sm font-medium text-gray-700">
          {LabelName}
        </label>
        {isSurpriseMe && (
          <button
            typeof='button'
            className="text-semibold text-xs bg-[#ECECF1] py-1 px-2 rounded-[5px] text-black"
          >
            Surprise Me
          </button>)}
      </div>
    </div>
  )
}

export default FormField
