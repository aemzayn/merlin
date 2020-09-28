# Copyright 2020 The Merlin Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# coding: utf-8

"""
    Merlin

    API Guide for accessing Merlin's model deployment functionalities  # noqa: E501

    OpenAPI spec version: 0.6.0
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six


class PredictionJobConfigModel(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'type': 'str',
        'uri': 'str',
        'result': 'PredictionJobConfigModelResult',
        'options': 'dict(str, str)'
    }

    attribute_map = {
        'type': 'type',
        'uri': 'uri',
        'result': 'result',
        'options': 'options'
    }

    def __init__(self, type='INVALID_MODEL_TYPE', uri=None, result=None, options=None):  # noqa: E501
        """PredictionJobConfigModel - a model defined in Swagger"""  # noqa: E501

        self._type = None
        self._uri = None
        self._result = None
        self._options = None
        self.discriminator = None

        if type is not None:
            self.type = type
        if uri is not None:
            self.uri = uri
        if result is not None:
            self.result = result
        if options is not None:
            self.options = options

    @property
    def type(self):
        """Gets the type of this PredictionJobConfigModel.  # noqa: E501


        :return: The type of this PredictionJobConfigModel.  # noqa: E501
        :rtype: str
        """
        return self._type

    @type.setter
    def type(self, type):
        """Sets the type of this PredictionJobConfigModel.


        :param type: The type of this PredictionJobConfigModel.  # noqa: E501
        :type: str
        """
        allowed_values = ["INVALID_MODEL_TYPE", "XGBOOST", "TENSORFLOW", "SKLEARN", "PYTORCH", "ONNX", "PYFUNC", "PYFUNC_V2"]  # noqa: E501
        if type not in allowed_values:
            raise ValueError(
                "Invalid value for `type` ({0}), must be one of {1}"  # noqa: E501
                .format(type, allowed_values)
            )

        self._type = type

    @property
    def uri(self):
        """Gets the uri of this PredictionJobConfigModel.  # noqa: E501


        :return: The uri of this PredictionJobConfigModel.  # noqa: E501
        :rtype: str
        """
        return self._uri

    @uri.setter
    def uri(self, uri):
        """Sets the uri of this PredictionJobConfigModel.


        :param uri: The uri of this PredictionJobConfigModel.  # noqa: E501
        :type: str
        """

        self._uri = uri

    @property
    def result(self):
        """Gets the result of this PredictionJobConfigModel.  # noqa: E501


        :return: The result of this PredictionJobConfigModel.  # noqa: E501
        :rtype: PredictionJobConfigModelResult
        """
        return self._result

    @result.setter
    def result(self, result):
        """Sets the result of this PredictionJobConfigModel.


        :param result: The result of this PredictionJobConfigModel.  # noqa: E501
        :type: PredictionJobConfigModelResult
        """

        self._result = result

    @property
    def options(self):
        """Gets the options of this PredictionJobConfigModel.  # noqa: E501


        :return: The options of this PredictionJobConfigModel.  # noqa: E501
        :rtype: dict(str, str)
        """
        return self._options

    @options.setter
    def options(self, options):
        """Sets the options of this PredictionJobConfigModel.


        :param options: The options of this PredictionJobConfigModel.  # noqa: E501
        :type: dict(str, str)
        """

        self._options = options

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(PredictionJobConfigModel, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, PredictionJobConfigModel):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
