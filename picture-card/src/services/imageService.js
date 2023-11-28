import axios from 'axios';

const baseUrl = `https://gist.githubusercontent.com/knot-freshket`;
const httpClient = axios.create();

export const getImages = async () => {
    let url = `${baseUrl}/142c21c3e8e54ef36e33f5dc6cf54077/raw/460fa6bd2bcc3aad83afde7256f1d742811f3392/freshket-places.json`;
    try {
        const res = await httpClient.get(url)
        return res.data;
    } catch (error) {
        return { error: true, ...error };
    }
}
export const getImageTags = async () => {
    let url = `${baseUrl}/fa49e0a5c6100d50db781f28486324d2/raw/55bc966f54423dc73384b860a305e1b67e0bfd7d/freshket-tags.json`;
    try {
        const res = await httpClient.get(url)
        return res.data;
    } catch (error) {
        return { error: true, ...error };
    }
}