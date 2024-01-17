from flask import Flask, request, send_file
import cv2
import numpy as np
import openvino.runtime as ov
import io
from PIL import Image

app = Flask(__name__)

# Load OpenVINO model
model_path = 'semantic-segmentation-adas-0001.xml'  # Update the path
core = ov.Core()
compiled_model = core.compile_model(model_path, 'CPU')

@app.route('/upload', methods=['POST'])
def upload_image():
    if 'file' not in request.files:
        return 'No file part', 400
    file = request.files['file']
    if file.filename == '':
        return 'No selected file', 400

    image = np.fromstring(file.read(), np.uint8)
    image = cv2.imdecode(image, cv2.IMREAD_COLOR)

    # Apply image processing operations
    blurred_image = cv2.GaussianBlur(image, (15, 15), 0)
    edges_image = cv2.Canny(image, 100, 200)
    sepia_filter = np.array([[0.272, 0.534, 0.131], [0.349, 0.686, 0.168], [0.393, 0.769, 0.189]])
    sepia_image = cv2.transform(image, sepia_filter)

    # Save or encode processed images and return them
    # Example: Saving and sending one image
    cv2.imwrite('blurred_image.jpg', blurred_image)
    return send_file('blurred_image.jpg', mimetype='image/jpeg')

# Add more routes for other functionalities as needed

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
