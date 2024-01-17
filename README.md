##Flask Code
image processing server that performs operations on images uploaded by users. Here's a breakdown of its functionality:

Imports and Flask App Initialization:

Various libraries are imported: Flask for the web server, cv2 (OpenCV) for image processing, numpy for numerical operations, openvino.runtime for potential model inference (although not used in the provided code), io, and PIL.Image for image manipulation.
A Flask app instance is created.
Load OpenVINO Model:

An OpenVINO model is loaded from the specified path semantic-segmentation-adas-0001.xml, but this model is not used in the code snippet provided.
Route /upload:

A route /upload is defined to handle POST requests for uploading images.
The code checks if the request contains a file part and returns an error message if not.
It then reads the uploaded file, ensuring there is a file selected and uploaded.
The image from the uploaded file is read into a NumPy array and decoded into a color image using OpenCV.
Image Processing:

The code applies various image processing operations to the uploaded image:
GaussianBlur: Applies a Gaussian blur with a kernel size of (15, 15).
Canny: Detects edges in the image using the Canny edge detection algorithm.
Sepia Filter: Applies a sepia color transformation to the image. However, this processed image is not used in any response.
Note: Although these operations are performed, only the blurred image is further utilized in this snippet.
Saving and Sending Processed Image:

The blurred image is saved as 'blurred_image.jpg'.
The saved blurred image is then sent back to the client as the response to the POST request.
Running the App:

The Flask application is configured to run on host='0.0.0.0' and port=5000, making it accessible over the network.

## Go Client

command-line application that uploads an image to a Flask server and saves the processed image returned by the server. Here's a step-by-step breakdown of its functionality:

Initialization:

The code sets up two important strings: serverURL, which is the URL of the Flask server's /upload endpoint, and filePath, the path to the image that needs to be uploaded. An outputFilePath is also defined to specify where the processed image should be saved.
Opening the Image File:

The program attempts to open the image file specified by filePath. If it fails to open the file (for example, if the file doesn't exist or there are permission issues), it prints an error message and exits.
Preparing a Multipart Form Request:

The program creates a new multipart form writer to build the request body.
It adds the image file to the form under the key "file". This matches the key that the Flask server expects for the uploaded file.

Sending the Request to the Flask Server:

A POST request is constructed with the multipart form data as the body. The Content-Type header is set to the appropriate multipart form data type.
The request is sent to the Flask server's /upload endpoint.
Handling the Server Response:

On receiving a response from the server, the program checks if the status code is http.StatusOK (HTTP 200). If it's not, it prints an error message indicating a non-OK status from the server.
If the response is OK, it creates a new file specified by outputFilePath to save the processed image.
The body of the response, which should be the processed image, is written to this file.
Final Output:

If everything is successful, the program saves the processed image and prints a message indicating the file where the processed image is saved.
In summary, this Go client application is designed to interact with the Flask server by uploading an image and saving the image processed by the server. This is useful in scenarios where image processing is done server-side, and clients need to send and receive images for processing.

## Complete Workflow


The complete workflow involving both the Flask (Python) backend and the Go client frontend is designed for uploading, processing, and retrieving images. Here's how the entire process works:

Flask Backend:
Server Initialization: The Flask app starts and listens for incoming requests on port 5000.

Image Upload Endpoint (/upload):

The app defines an endpoint /upload to handle POST requests for image uploads.
When a request is received, it checks for an image file in the request. If no file is found or if there is an issue with the file, it returns an error.
If a valid image file is provided, the Flask app performs several image processing operations on the image using OpenCV, like Gaussian blurring.
The processed image (in this case, the blurred image) is saved as a JPEG file (blurred_image.jpg).
Response:

The Flask app sends the processed image file back to the client as a response to the POST request.
Go Client:
User Input: The Go client is started, and it is configured with the path to the image file to be uploaded (1.jpg).

Preparing and Sending Request:

The Go application prepares a POST request with the image file in a multipart form.
It sends this request to the Flask server's /upload endpoint.
Receiving and Saving the Processed Image:

The client receives the response from the Flask server, which contains the processed image.
If the response is successful, the Go client saves the received image to the specified output file path (2.jpg).
Complete Workflow:
Starting the Process:

The user runs the Go client and specifies the image file to be processed.
Uploading Image to Flask Server:

The Go client sends the image file to the Flask server for processing.
Server-Side Image Processing:

The Flask server receives the image, performs specified processing operations (e.g., blurring), and saves the processed image to a file.
Receiving Processed Image:

The Flask server sends the processed image back to the Go client.
Saving the Processed Image:

The Go client saves the processed image received from the server to the local filesystem.
Completion:

The user is notified that the processed image has been saved, completing the cycle of uploading, processing, and retrieving an image.
This workflow effectively demonstrates a client-server architecture where the heavy lifting of image processing is handled by the server, and the client manages the user interactions and handles sending/receiving data.
