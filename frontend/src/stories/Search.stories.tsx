import React, { useState } from 'react';
import Search from '../components/Search';

export default {
  title: 'Components/Search',
  component: Search,
};

export const Default = () => <Search></Search>;
export const SeeInput = () => {
  const [input, setInput] = useState('Test');

  return (
    <div>
      <Search
        value={input}
        onInput={(e) =>
          setInput((e.target as HTMLInputElement).value)
        }></Search>
      <p className="mt-2 text-gray-400">Input Data: {input}</p>
    </div>
  );
};
