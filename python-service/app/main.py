from typing import List

import cv2
import numpy as np
from fastapi import FastAPI, File, HTTPException, UploadFile
from pydantic import BaseModel

app = FastAPI(title="GradeFlowOCR")


class GabaritoResult(BaseModel):
    answers: List[str]
    result: int


@app.post("/process-gabarito", response_model=GabaritoResult)
async def process_gabarito(file: UploadFile = File(...)):
    with open("temp_image.jpg", "wb") as f:
        f.write(await file.read())


    image = cv2.imread("temp_image.jpg")
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)
    blurred = cv2.GaussianBlur(gray, (5, 5), 0)
    _, binary = cv2.threshold(blurred, 100, 255, cv2.THRESH_BINARY_INV)

    circles = cv2.HoughCircles(binary, cv2.HOUGH_GRADIENT, dp=1.2, minDist=20,
                               param1=50, param2=30, minRadius=10, maxRadius=20)

    answers = []

    if circles is not None:
        circles = np.round(circles[0, :]).astype("int")

        for (x, y, r) in circles:
            circle_region = binary[y - r:y + r, x - r:x + r]

            filled_ratio = np.sum(circle_region == 255) / (np.pi * r * r)

            is_marked = filled_ratio > 0.5

            if is_marked:
                answers.append(f"({x}, {y})")

            color = (0, 255, 0) if is_marked else (0, 0, 255)
            cv2.circle(image, (x, y), r, color, 2)

    # Display the processed image (for debugging purposes)
    cv2.imshow("Processed Answer Sheet", image)
    cv2.waitKey(0)
    cv2.destroyAllWindows()

    return GabaritoResult(
        answers=answers,
        result=len(answers)
    )

###TODO###
def process_gabarito(image_path: str) -> GabaritoResult:
    image = cv2.imread(image_path)
    gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)

    blurred = cv2.GaussianBlur(gray, (9, 9), 0)

        pass

    circles = cv2.HoughCircles(
        blurred,
        cv2.HOUGH_GRADIENT,
        dp=1.2,
        minDist=20,
        param1=50,
        param2=30,
        minRadius=10,
        maxRadius=20
    )

    if circles is not None:
        circles = np.round(circles[0, :]).astype("int")

        for (x, y, r) in circles:
            roi = gray[y - r:y + r, x - r:x + r]

            _, thresh = cv2.threshold(roi, 100, 255, cv2.THRESH_BINARY_INV)

            non_zero = cv2.countNonZero(thresh)
            area = np.pi * (r ** 2)
            fill_ratio = non_zero / area

            if fill_ratio > 0.35:
                cv2.circle(image, (x, y), r, (0, 255, 0), 3)
            else:
                cv2.circle(image, (x, y), r, (0, 0, 255), 3)

    # Exibir a imagem processada
    cv2.imshow("Gabarito Processado", image)
    cv2.waitKey(0)
    cv2.destroyAllWindows()

