import React, { SyntheticEvent } from 'react';
import Icon from './Icon';

interface SearchProps {
  value?: string;
  onInput?: (event: SyntheticEvent) => void;
}

const Search: React.FC<SearchProps> = (props): JSX.Element => {
  return (
    <div className="relative w-full max-w-xs 2xl:max-w-lg">
      <Icon
        name="search"
        className="absolute inset-y-0 left-0 w-6 h-6 my-auto ml-2 text-gray-300"
      />
      <input
        placeholder="Search"
        type="text"
        className="w-full py-2 pl-10 pr-4 text-sm text-gray-700 placeholder-gray-300 transition duration-300 ease-in-out bg-gray-100 rounded-lg focus:outline-none focus:shadow-md"
        {...props}
      />
    </div>
  );
};

export default Search;
