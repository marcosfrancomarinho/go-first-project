import React from 'react';

export const useConfirmDeleteProduct = () => {
  const [showBox, setShowBox] = React.useState<boolean>(false);
  const [id, setId] = React.useState<string>('');
  const [index, setIndex] = React.useState<number>(0);

  const choiseProductToRemove = React.useCallback(
    async (id: string, index: number) => {
      setShowBox(true);
      setId(id);
      setIndex(index);
    },
    [index, showBox, index]
  );

  return { id, index, showBox, choiseProductToRemove , setShowBox};
};
