import React, { useEffect, useRef, useState } from 'react';

interface IconProps extends React.SVGProps<HTMLDivElement> {
  name: string;
}

const Icon: React.FC<IconProps> = ({ name, ...rest }): JSX.Element | null => {
  const ImportedIconRef = useRef<React.FC<React.SVGProps<SVGSVGElement>>>();
  const [loading, setLoading] = useState(false);

  useEffect((): void => {
    setLoading(true);
    const importIcon = async (): Promise<void> => {
      try {
        ImportedIconRef.current = (
          await import(`../../assets/icons/${name}.svg?component`)
        ).default;
      } catch (error) {
        throw error;
      } finally {
        setLoading(false);
      }
    };

    importIcon();
  }, [name]);

  if (!loading && ImportedIconRef.current) {
    const { current: ImportedIcon }: any = ImportedIconRef;
    return (
      <div {...rest}>
        <ImportedIcon />
      </div>
    );
  }

  return null;
};

export default Icon;
