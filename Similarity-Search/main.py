import onnxruntime as ort

session = ort.InferenceSession("model.onnx", providers=["CPUExecutionProvider"])

input_name = session.get_inputs()[0].name
output_input = session.get_outputs()[0].name


