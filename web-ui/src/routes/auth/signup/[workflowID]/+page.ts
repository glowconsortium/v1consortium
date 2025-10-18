/** @type {import('./$types').PageLoad} */
export function load({ params }) {
  // `params` is an object that contains the dynamic parts of the URL.
  // The property name corresponds to the folder name, in this case, `workflowid`.
  return {
    workflowID: params.workflowID
  };
}
