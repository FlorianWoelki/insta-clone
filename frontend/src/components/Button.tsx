import React, { FunctionComponent } from 'react';

type ButtonProps = {};

const Button: FunctionComponent<ButtonProps> = (props): JSX.Element => {
  return (
    <button
      type="button"
      className="flex items-center px-4 py-2 space-x-1 font-medium text-white rounded-lg bg-gradient-to-r from-pink-600 to-yellow-500 hover:from-pink-800 hover:to-yellow-700 focus:outline-none">
      {props.children}
    </button>
  );
};

export default Button;
