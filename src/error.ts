export interface NodeErrorMaybe extends Error {
  code?: string;
}
/** Handle errors and conditionally exit program */
export const handleError = (error: Error | unknown) => {
  if (
    (error as NodeErrorMaybe)?.code !== 'ENOENT' &&
    (error as NodeErrorMaybe)?.code !== 'EEXIST'
  ) {
    console.error(error);
    process.exit(1);
  }
};
