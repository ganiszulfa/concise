import { local, prod } from '$lib/configs';

export async function GraphQLQuery(query, variables) {
    let base = local.api_csr
    if (import.meta.env.PROD) {
        base = prod.api_csr
    }

    var myHeaders = new Headers();
    myHeaders.append("Content-Type", "application/json");

    var graphql = JSON.stringify({
        query: query,
        variables: variables,
    });
    var requestOptions = {
        method: 'POST',
        headers: myHeaders,
        body: graphql,
        redirect: 'follow'
    };

    var Data = null;

    await fetch(base + "/graphql", requestOptions)
        .then(response => response.text())
        .then(result => {
            Data = JSON.parse(result).data;
        })
        .catch(error => console.log('error', error));
    return Data;
}