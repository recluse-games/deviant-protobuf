// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: Board.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Deviant {

  /// <summary>Holder for reflection information generated from Board.proto</summary>
  public static partial class BoardReflection {

    #region Descriptor
    /// <summary>File descriptor for Board.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static BoardReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CgtCb2FyZC5wcm90bxIHRGV2aWFudBoLVGlsZXMucHJvdG8aDkVudGl0aWVz",
            "LnByb3RvIksKBUJvYXJkEh0KBXRpbGVzGAEgASgLMg4uRGV2aWFudC5UaWxl",
            "cxIjCghlbnRpdGllcxgCIAEoCzIRLkRldmlhbnQuRW50aXRpZXNCCVoHZGV2",
            "aWFudGIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Deviant.TilesReflection.Descriptor, global::Deviant.EntitiesReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Deviant.Board), global::Deviant.Board.Parser, new[]{ "Tiles", "Entities" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class Board : pb::IMessage<Board> {
    private static readonly pb::MessageParser<Board> _parser = new pb::MessageParser<Board>(() => new Board());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<Board> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Deviant.BoardReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Board() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Board(Board other) : this() {
      tiles_ = other.tiles_ != null ? other.tiles_.Clone() : null;
      entities_ = other.entities_ != null ? other.entities_.Clone() : null;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Board Clone() {
      return new Board(this);
    }

    /// <summary>Field number for the "tiles" field.</summary>
    public const int TilesFieldNumber = 1;
    private global::Deviant.Tiles tiles_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Deviant.Tiles Tiles {
      get { return tiles_; }
      set {
        tiles_ = value;
      }
    }

    /// <summary>Field number for the "entities" field.</summary>
    public const int EntitiesFieldNumber = 2;
    private global::Deviant.Entities entities_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Deviant.Entities Entities {
      get { return entities_; }
      set {
        entities_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as Board);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(Board other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (!object.Equals(Tiles, other.Tiles)) return false;
      if (!object.Equals(Entities, other.Entities)) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (tiles_ != null) hash ^= Tiles.GetHashCode();
      if (entities_ != null) hash ^= Entities.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (tiles_ != null) {
        output.WriteRawTag(10);
        output.WriteMessage(Tiles);
      }
      if (entities_ != null) {
        output.WriteRawTag(18);
        output.WriteMessage(Entities);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (tiles_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Tiles);
      }
      if (entities_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Entities);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(Board other) {
      if (other == null) {
        return;
      }
      if (other.tiles_ != null) {
        if (tiles_ == null) {
          Tiles = new global::Deviant.Tiles();
        }
        Tiles.MergeFrom(other.Tiles);
      }
      if (other.entities_ != null) {
        if (entities_ == null) {
          Entities = new global::Deviant.Entities();
        }
        Entities.MergeFrom(other.Entities);
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            if (tiles_ == null) {
              Tiles = new global::Deviant.Tiles();
            }
            input.ReadMessage(Tiles);
            break;
          }
          case 18: {
            if (entities_ == null) {
              Entities = new global::Deviant.Entities();
            }
            input.ReadMessage(Entities);
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code
